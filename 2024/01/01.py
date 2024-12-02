def absolute_value(x):
    return x if x >= 0 else -x

def measure_distance(first, second):
    distance = 0
    first.sort()
    second.sort()
    for i in range(0, len(first), 1):
        distance += absolute_value(first[i] - second[i])
    print(f"distance: {distance}")

def measure_similarity(first, second):
    similarity_score = 0
    for i in first:
        similarity_score += i * second.count(i)
    print(f"similarity score: {similarity_score}")

def get_lists():
    first, second = [], []
    with open("input.txt", "r") as f:
        for line in f:
            parts = line.split()
            if len(parts) == 2:
                first.append(int(parts[0]))
                second.append(int(parts[1]))
    return first, second

if __name__ == "__main__":
    first, second = get_lists()
    measure_distance(first, second)
    measure_similarity(first, second)