import random

def bubble_sort(l: list) -> list:
    for i in range(len(l) - 1):
        for j in range(len(l) - (i + 1)):
            if l[j] > l[j+1]:
                l[j], l[j+1] = l[j+1], l[j]
    return l


if __name__ == "__main__":
  x = [random.randint(0, 100) for _ in range(10)]
  print(f"=> {bubble_sort(x)}")