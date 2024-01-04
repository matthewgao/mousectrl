import pyautogui
import time, json
from flask import Flask, request, make_response, Response
from threading import Thread, Lock
pyautogui.PAUSE = 0.01
pyautogui.FAILSAFE = False
mutex = Lock()
# def move_mouse(x, y):
#     pyautogui.moveTo(x, y)

# def main():
#     print("Move the mouse to (100, 100)")
#     move_mouse(100, 100)
#     time.sleep(2)

#     print("Click at the current mouse position")
#     pyautogui.click()
#     time.sleep(2)

#     print("Move the mouse to (300, 300)")
#     move_mouse(300, 300)
#     time.sleep(2)

#     pyautogui.dragTo(400, 550, button="left")

app = Flask(__name__)

@app.route('/moveMouse', methods=['POST'])
def moveMouse():
    body = request.get_data()
    # print(body)
    body_json = json.loads(body)
    with mutex:
        for value in body_json:
            x = value['x']
            y = value['y']
            is_drag = value['is_drag']
            if is_drag:
                # print(x)
                pyautogui.dragTo(x, y, button="left")
            else:
                pyautogui.moveTo(x, y)

    return make_response("ok", 200)

if __name__ == '__main__':
    app.debug = True
    app.config["JSON_AS_ASCII"] = False
    app.run(host='0.0.0.0', port=8000)