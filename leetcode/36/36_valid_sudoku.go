package main

func isValidSudoku(board [][]byte) bool {
    rows := [9][9]bool{}
    cols := [9][9]bool{}
    diags := [2][9]bool{}
    boxs := [9][9]bool{}
    for i:=0;i<9;i++ {
        for j:=0;j<9;j++ {
            b := board[i][j]
            if b == '.' {
                continue
            }
            n := b - '1'
            if rows[i][n] {
                return false
            }
            if cols[j][n] {
                return false
            }
            if i == j {
                if diags[0][n] {
                    return false
                }
                diags[0][n] = true
            }
            if i+j == 8 {
                if diags[1][n] {
                    return false
                }
                diags[1][n] = true
            }
            boxIdx := (i / 3) * 3 + j / 3
            if boxs[boxIdx][n] {
                return false
            }
            boxs[boxIdx][n] = true
            rows[i][n] = true
            cols[j][n] = true
        }
    }
    return true
}

func main() {

}