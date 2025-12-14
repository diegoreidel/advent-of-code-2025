from z3 import *


def process_line(line):
    opt = Optimize()

    parts = line.split()
    joltages = parts.pop()[1:-1].split(',')
    buttons = parts[1:]
    constraints = {}

    z3_vars = [Int(f'button_{i}') for i, _ in enumerate(buttons)]

    for var in z3_vars:
        opt.add(var >= 0)

    for i, button in enumerate(buttons):
        for pos in button[1:-1].split(','):
            if int(pos) not in constraints:
                constraints[int(pos)] = []

            constraints.get(int(pos)).append(z3_vars[i])

    for i, goal in enumerate(joltages):
        opt.add(Sum(constraints.get(i)) == goal)

    l = opt.minimize(Sum(z3_vars))

    if opt.check() == sat:
        return opt.lower(l).as_long()
    else:
        print("Problem is unsatisfiable")

    return 0

# Open the file using a context manager, which ensures it is closed automatically
with open('input/input.txt', 'r') as file:

    answer = 0

    for line in file:
        answer += process_line(line.strip())

    print("Final Answer:", answer)