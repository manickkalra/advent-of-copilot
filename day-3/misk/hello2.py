# # chat gpt wrote below 
# def find_priority(item_type):
#     if item_type.islower():
#         return ord(item_type) - ord('a') + 1
#     else:
#         return ord(item_type) - ord('A') + 27
# # chat gpt wrote below
# def solve(rucksack_contents):
#     priority_sum = 0
#     for rucksack in rucksack_contents:
#         first_compartment = set(rucksack[:len(rucksack)//2])
#         second_compartment = set(rucksack[len(rucksack)//2:])
#         common_items = first_compartment.intersection(second_compartment)
#         for item in common_items:
#             priority_sum += find_priority(item)
#     return priority_sum

# day 2 chatgpt updated below
def find_priority(item_type):
    if item_type.islower():
        return ord(item_type) - ord('a') + 1
    else:
        return ord(item_type) - ord('A') + 27

def solve(rucksack_contents):
    priority_sum = 0
    for i in range(0, len(rucksack_contents), 3):
        group_badge = set(rucksack_contents[i])
        for j in range(1, 3):
            group_badge.intersection_update(set(rucksack_contents[i+j]))
        for badge in group_badge:
            priority_sum += find_priority(badge)
    return priority_sum

# immediately prompted by copilot for below
def main():
    rucksack_contents = []
    with open('input.txt') as f:
        for line in f:
            # print(line)
            rucksack_contents.append(line.strip())
    print(solve(rucksack_contents))
# I wrote two character of code "if", copilot did the rest
if __name__ == '__main__':
    main()