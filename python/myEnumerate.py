def myEnumerate(iterable, start=0):
    it = iter(iterable)
    while True:
        try:
            yield start, next(it)
            start += 1
        except StopIteration:
            del it
            return


if __name__ == '__main__':
    a = [i for i in range(1, 15)]
    for idx, ai in myEnumerate(a, 2):
        print(f"{idx}:{ai}")
