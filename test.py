import requests
import json
from datetime import datetime,timezone

BASE_URL = "http://106.15.77.79:8089/CreateTask"  # 请根据实际情况修改端口和地址

# 插入测试
# 插入测试
def test_insert_task():
    current_timestamp = int(datetime.now(timezone.utc).timestamp() * 1000)  # 获取当前毫秒级时间戳
    print("current_timestampcurrent_timestampcurrent_timestamp",current_timestamp)
    task_data = {
        "time": current_timestamp,  # 将时间戳转换为字符串
        "phonenumber": "13059088401",
        "jobstatus": -1,
        "jobinfo": json.dumps({
            "table_number": 4,
            "table_info": [{"type": "红烧肉", "number": "4"},{"type":"米饭","number":"2"}]
        })
    }

    print(task_data)

    response = requests.post(BASE_URL, json=task_data)
    print("Insert Task Response:", response.json())

# 查询测试
def test_get_tasks():
    response = requests.get(BASE_URL)
    print("Get Tasks Response:",response.json())
    print("Get Tasks Response--------:",len(response.json()))
    for k ,v in response.json()[0].items():
        print(k,v)

def test_get_task_by_time_and_user(task_time, phonenumber):
    response = requests.get(BASE_URL + "/search", params={"time": task_time, "phonenumber": phonenumber})
    print(BASE_URL + "/search")
    print("Get Task by Time and User Response:", response.text)

# 更新测试
def test_update_task(task_time):
    update_data = {
        "phonenumber": "new_phonenumber",
        "jobstatus": 1,
        "jobinfo": json.dumps({
            "table_number": 5,
            "table_info": [[{"type": "红烧肉", "number": "5"}]]
        })
    }


    response = requests.put(f"{BASE_URL}/{task_time}", json=update_data)
    print("Update Task Response:", response.json())

# 删除测试
def test_delete_task(task_time):
    response = requests.delete(f"{BASE_URL}/{task_time}")
    print("Delete Task Response:", response.json())

if __name__ == "__main__":
    # 先插入一个任务
    test_insert_task()

    # 查询任务列表
    test_get_tasks()

      # 假设已经插入的时间和用户名
    task_time_to_query = "1730132217832"  # 请替换为实际插入的时间
    phonenumber_to_query = "13059088401"  # 替换为实际插入的用户名

    # 根据时间戳和用户名查询任务
    test_get_task_by_time_and_user(task_time_to_query, phonenumber_to_query)

    # # 更新任务（请根据实际插入的时间修改此处）
    # task_time_to_update = datetime.now().isoformat()  # 此处需要实际插入的
