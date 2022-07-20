#!/usr/bin/env python3

import time


def is_white_space(c):
    return c == '\r' or c == '\n' or c == ' ' or c == '\t'


def word_count(file_path: str):
    with open(file_path, "rt") as file:
        str = file.read()

        wc = 0
        is_prev_white_space = True
        for i in range(1, len(str)):
            is_cur_white_space = is_white_space(str[i])

            if is_cur_white_space and not is_prev_white_space:
                wc += 1
            is_prev_white_space = is_cur_white_space

        if len(str) != 0 and not is_prev_white_space:
            wc += 1

        print(f'Num Words = {wc}\n')


start = time.time()
word_count("file.txt")
end = time.time()
elapsed = (end - start) * 1000
print(f"Elapsed = {elapsed} ms\n")
