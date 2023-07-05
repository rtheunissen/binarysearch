set xlabel "Δ"
set ylabel "Γ"

set format x "%.2f"
set format y "%.2f"

set grid ytics back
set grid xtics back

set border 3 front lc '#000000' lt 1 lw 0

set xrange [0:5]
set yrange [0:5]

set object circle at (3.),(2.) size 0.05
set object circle at (3.),(4./3.) size 0.05
set object circle at (2.),(3./2.) size 0.05
set object circle at (1.+sqrt(2)),(sqrt(2)) size 0.05

plot "wb_topdown_polytope.csv" using 1:2 with dots lc rgb "#000000" notitle

