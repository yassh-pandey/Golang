// Source: https://leetcode.com/problems/island-perimeter/
func islandPerimeter(grid [][]int) int {
    var parameter int
    var maxRows = len(grid)
    var maxColumns = 0
    if maxRows > 0{
        maxColumns = len(grid[0])
    }
    if maxRows == 0 || maxColumns == 0{
        return 0
    }
    for i, row := range grid{
        for j, tile := range row{
            if tile == 1{
                // We got ourself into land
                // Check corners and compute contribution to parameter
                
                // Check left
                if j == 0  || ( j-1 >= 0 && grid[i][j-1] == 0 ){
                    parameter+=1
                }
                
                // Check right
                if ( j == maxColumns - 1 ) || ( j+1 < maxColumns && grid[i][j+1] == 0){
                    parameter+=1
                }
                
                // Check top
                if i == 0 || ( i-1 >= 0 && grid[i-1][j] == 0){
                    parameter+=1
                } 
                
                // Check bottom
                if i == maxRows -1 || (i + 1 < maxRows && grid[i+1][j] == 0){
                    parameter+=1
                } 
            }
        }
    }
    return parameter
}
