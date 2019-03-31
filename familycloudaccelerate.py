import requests
import hmac
import datetime
import time
import sys
import json

ACCESS_URL = "/family/qos/startQos.action"

UP_QOS_URL = "http://api.cloud.189.cn/family/qos/startQos.action"


def mac_sha1(params, secret):
    key = bytes(secret, encoding='utf-8')
    msg = bytes(params, encoding='utf-8')
    hash_mac = hmac.new(key=key, msg=msg, digestmod='SHA1')
    return hash_mac.hexdigest().upper()


def get_signature(access_url, session_key, session_secret, request_method, date):
    """
    params='SessionKey={session_key}&Operate={request_method}&RequestURI={access_url}&Date={date}'
    :param access_url: str
    :param session_key: str
    :param session_secret: str
    :param request_method: str
    :param date: str
    :return: mac_sha1
    """

    params = "SessionKey={}&Operate={}&RequestURI={}&Date={}".format(
        session_key,
        request_method,
        access_url,
        date
    )
    return mac_sha1(params, session_secret)


def create_date():
    """
    :return: str 'Sun, 31 Mar 2019 05:35:33 GMT'
    """
    time_stop = 16000/1000.0
    time_start = 12500/1000.0
    time_stamp = time.time()+(time_stop-time_start)
    time_utc = datetime.datetime.utcfromtimestamp(time_stamp)
    gmt_format = "%a, %d %b %Y %H:%M:%S GMT"
    return time_utc.strftime(gmt_format)


def heart_beat(session, secret, method="GET", **kwargs):
    extra_header = kwargs.get("extra_header")
    send_data = kwargs.get("send_data")
    date = create_date()
    signature = get_signature(ACCESS_URL, session, secret, method, date)
    print("heart_beat:<signature:{}>".format(signature))
    print("date:<{}>".format(date))
    header = {
        "SessionKey": session,
        "Signature": signature,
        "Date": date
    }
    if extra_header:
        header.update(extra_header)
    if method == "POST":
        resp = requests.get(UP_QOS_URL, params=send_data, headers=header)
        return resp
    return requests.post(UP_QOS_URL, data=send_data, headers=header)


if __name__ == "__main__":
    config = None
    if len(sys.argv) >= 2:
        config_file = open(sys.argv[1], mode="r")
        config = json.load(config_file)
    else:
        config_file = open("./config.json", mode="r" )
        config = json.load(config_file)
    session_key = config["session_key"]
    session_secret = config["session_secret"]
    setting = config["setting"]
    method = setting['method']
    rate = setting['rate']
    extra_header = config["extra_header"]
    send_data = config["send_data"]

    count = 0

    while True:
        count += 1
        print("Sending heart_beat package <{}>".format(count))
        resp = heart_beat(session_key, session_secret,
                          method=method,
                          extra_header=extra_header,
                          send_data=send_data
                          )
        print("status_code:{}".format(resp.status_code))
        print("response:\n{}".format(resp.text))
        print("Send heart_beat <{}> package Success".format(count))
        print("*******************************************")
        time.sleep(rate)

