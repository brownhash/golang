import boto3
from datetime import datetime, timedelta
from flask import Flask, render_template
from flask import request as rq


app = Flask(__name__)

@app.route("/metrics-mumbai")
def metrics_mumbai():
    return metrics("ap-south-1", 'i-005e8a7e3c38bab73')

@app.route("/metrics-singapore")
def metrics_sg():
    return metrics("ap-southeast-1", 'i-03be561ef6bc83400')

@app.route("/get-metrics-sg")
def get_metrics():
    reader = open('./templates/data', "r")
    value = reader.read()
    reader.close()
    client = boto3.client('ec2', region_name='ap-southeast-1')
    request = client.describe_instances().get('Reservations')

    id = ""
    for groups in request:
        for instances in groups.get('Instances'):
            tags = instances.get('Tags')
            for i in tags:
                if i.get('Key') == 'Name' and i.get('Value') == value:
                    id = instances.get('InstanceId')

    if len(id)>0:
        data = metrics('ap-southeast-1', id)
        return data

    else:
        return render_template("instance_error.html")


def metrics(region, i_id):
    t_day = datetime.today().day
    t_month = datetime.today().month
    t_year = datetime.today().year

    client = boto3.client('cloudwatch', region_name=region)
    response = client.get_metric_statistics(
        Namespace='AWS/EC2',
        MetricName='CPUUtilization',
        Dimensions=[
            {
                'Name': 'InstanceId',
                'Value': i_id
            },
        ],
        StartTime=datetime(2019, 9, 1) - timedelta(seconds=300),
        EndTime=datetime(t_year, t_month, t_day),
        Period=86400,
        Statistics=[
            'Maximum',
            'Minimum',
            'Average'
        ],
        Unit='Percent'
    )
    metric = ""
    for cpu in response['Datapoints']:
        date = str(cpu.get('Timestamp')).split(" ")[0]
        date = date.split("-")
        date = date[2]+"-"+date[1]+"-"+date[0]
        data = "{} {} {} {}$$".format(date, int(cpu.get('Minimum')),
                                      int(cpu.get('Maximum')), int(cpu.get('Average')))
        metric += data
    max="31-12-"+str(datetime.today().year)

    def compare_date(date1, date2):
        date1 = date1.split("-")
        d1 = int(date1[0])
        m1 = int(date1[1])
        y1 = int(date1[2])

        date2 = date2.split("-")
        d2 = int(date2[0])
        m2 = int(date2[1])
        y2 = int(date2[2])

        if y2 < y1:
            return True
        else:
            if m2 < m1:
                return True
            else:
                if d2 < d1:
                    return True
                else:
                    return False

    split_data = metric.split("$$")
    return_data = ""
    min = split_data[1].split(" ")[0]
    for i in range(0, len(split_data)-1):
        for j in range(0, len(split_data)-1):
            if compare_date(split_data[i].split(" ")[0], split_data[j].split(" ")[0]):
                temp = split_data[j]
                split_data[j] = split_data [i]
                split_data[i] = temp
    for i in split_data:
        return_data = return_data + i + "$$"
    return return_data

if __name__ == "__main__":
    app.run(debug=True)