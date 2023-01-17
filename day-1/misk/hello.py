#  chat-gpt for all unless otherwise specified
# used chat gpt to for day 1 part 1 advent of code challenge to find elf carrying the most calories
from pathlib import Path
input_str = Path('/Users/manickkalra/Projects/advent-of-code/advent-of-copilot/day-1/misk/input.txt').read_text()
input_list = input_str.strip().split('\n\n')
elf_food = [[int(x) for x in item.split('\n')] for item in input_list]
max_calories = 0
max_elf = None
top_3_elf_calories = [0,0,0]
for i, food_items in enumerate(elf_food):
    total_calories = sum(food_items)
    # manick did below
    lowest_calories = min(top_3_elf_calories)
    # copilot did below
    if total_calories > lowest_calories:
        top_3_elf_calories.remove(lowest_calories)
        top_3_elf_calories.append(total_calories)
        # break
    print(total_calories, top_3_elf_calories, sum(top_3_elf_calories))

print(f"Elf {max_elf} is carrying the most calories: {max_calories}")
print(f"Top 3 elves with the most calories: {sum(top_3_elf_calories)}")
