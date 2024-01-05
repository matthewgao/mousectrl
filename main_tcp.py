import pyautogui
import time, json
pyautogui.PAUSE = 0.0
# pyautogui.FAILSAFE = False
pyautogui.MINIMUM_DURATION = 0.0
pyautogui.MINIMUM_SLEEP = 0.0
pyautogui.DARWIN_CATCH_UP_TIME = 0.0

import socket
import struct
import threading

def handle_client(client_socket, client_address):
    print(f"Got a connection from {client_address}")
    try:
        x = 0;
        y = 0;
        drag = 0;
        has_x = False
        has_y = False
        has_drag = False
        
        while True:
            data = client_socket.recv(4)
            if not data:
                print(f"Client {client_address} disconnected.")
                break
            received_int = struct.unpack('>I', data)[0]
            # print(f"Received integer from {client_address}: {received_int}")
            if has_x == False:
                x = received_int
                has_x = True
            elif has_y == False:
                y = received_int
                has_y = True
            elif has_drag == False:
                drag = received_int
                has_drag = True

            if has_x == True and has_y == True and has_drag == True:
                if drag == 1:
                    # print(x)
                    pyautogui.dragTo(x, y, button="left")
                else:
                    pyautogui.moveTo(x, y)
                has_x = False
                has_y = False
                has_drag = False

    except Exception as e:
        print(f"Error with {client_address}: {e}")
    finally:
        client_socket.close()

def start_server(host, port):
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.bind((host, port))
    server_socket.listen(5)
    print(f"Server listening on {host}:{port}")

    try:
        while True:
            client_socket, client_address = server_socket.accept()
            client_thread = threading.Thread(target=handle_client, args=(client_socket, client_address))
            client_thread.daemon = True
            client_thread.start()
    finally:
        server_socket.close()

host = '0.0.0.0'
port = 8080
start_server(host, port)
