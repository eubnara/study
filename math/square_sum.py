
square = 1
sum = 1
sum_idx = 1
square_idx = 1

while True:
    if square == sum:
        print square
        print "square_idx: ", square_idx
        print "sum_idx: ", sum_idx
        ans = raw_input("continue? (y/n): ")
        if ans == "n":
            break
        else:
            square_idx += 1
            square = square_idx * square_idx
            sum_idx += 1
            sum += sum_idx
    else:
        if square > sum:
            sum_idx += 1
            sum += sum_idx
        else:
            square_idx += 1
            square = square_idx * square_idx
