def calculate_score(strategy_guide):
    move_map = {'A': 'X', 'B': 'Y', 'C': 'Z'}
    score_map = {'win': 6, 'lose': 0, 'draw': 3}
    score_value = {'X': 1, 'Y': 2, 'Z': 3}
    total_score = 0
    for moves in strategy_guide:
        # print(moves)
        opponent_move = moves[0]
        our_move = moves[1]
        if move_map[opponent_move] == our_move:
            outcome = 'draw'
        elif (opponent_move == 'A' and our_move == 'Z') or (opponent_move == 'B' and our_move == 'X') or (opponent_move == 'C' and our_move == 'Y'):
            outcome = 'lose'
        else:
            outcome = 'win'
        score = score_map[outcome] + score_value[our_move]
        total_score += score
    return total_score

# main loop
from pathlib import Path
if __name__ == '__main__':
    input_file = Path.cwd() / 'input.txt'
    input_str = input_file.read_text()
    input_str = input_str.split('\n')
    print(len(input_str))
    strategy_guide = []
    for line in (input_str):
        # print(line.strip())
        strategy_guide.append(line.split(' '))
    strategy_guide.pop()
    # print(strategy_guide)
    # print(input_str)
    # strategy_guide = [tuple(line.strip().split()) for line in input_str.split('\n')]
    print(calculate_score(strategy_guide))