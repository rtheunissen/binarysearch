set terminal svg size 400,300 dynamic

set xlabel "Δ" font ",14" offset 0,0
set ylabel "Γ" font ",14" offset 0,0 rotate by 0

set tmargin
set lmargin
set bmargin
set rmargin 6

set style line 1000 dashtype 3 lw 0.5 pt 1 ps 0 lc "#aaaaaa"

set ytics ("1" 1, "2" 2, "5/3" 5./3, "3/2" 1.5, "√2" sqrt(2), "4/3" 4./3., "5/4" 5./4)
set xtics ("2" 2, "1+√2" 1.+sqrt(2), "3" 3, "7/2" 3.5, "4/3" 4./3., "4" 4, "9/2" 4.5, "5" 5)

set y2tics ("α = 2/11" 11./9, "α = 1-√2/2" sqrt(2))
set x2tics ("α = 2/11"  9./2, "α = 1-√2/2" 1.+sqrt(2))

set grid ytics back ls 1000
set grid xtics back ls 1000

set grid y2tics back ls 1000
set grid x2tics back ls 1000

set border 3 front lc '#000000' lt 1 lw 0

#set style rectangle lc 0 fc rgb "#FFFF00" fs solid 0.2 noborder
#set obj rect from (9./2),(11./9) to (1.+sqrt(2)),(sqrt(2)) front

set xrange [1.8:5]
set yrange [0.9:2.1]
set y2range [0.9:2.1]

plot "wb_topdown_polytope_many.csv" using 1:2 with points axes x1y1 pt 5 ps 0.175 lc rgb "#000000" notitle, \
     "wb_topdown_polytope_many.csv" using 1:(($1+1)/$1) with lines axes x1y1 lc rgb "#00FFFF" dt 2 lw 1  notitle, \
     "wb_topdown_polytope_many.csv" using 1:2 with points axes x1y2 pt 5 ps 0.0 lc rgb "#000000" notitle

