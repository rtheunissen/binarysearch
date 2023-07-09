set terminal svg \
    size 400,200 \
    dynamic \
    font "docs/katex/fonts/KaTeX_Main-Regular.ttf,8"

set xlabel      "Δ"     offset 0,0
set ylabel      "Γ"     offset 0,0 rotate by 0

set lmargin 8
set bmargin 4
set rmargin 0
set tmargin 1

set style line 1000 dashtype 3 lw 0.1 pt 1 ps 0 lc black

set ytics ("1" 1, "2" 2, "5/3" 5./3, "3/2" 1.5, "√2" sqrt(2), "4/3" 4./3.)
set xtics ("2" 2, "1+√2" 1.+sqrt(2), "3" 3, "7/2" 3.5, "4/3" 4./3., "4" 4, "9/2" 4.5, "5" 5)

set grid ytics back ls 1000
set grid xtics back ls 1000

set border 3 front lt 1 lw 0

set xrange [1.8:5]
set yrange [0.9:2.1]
set y2range [0.9:2.1]

set label "Γ ≥ (Δ + 1) / Δ" at graph 0.85, graph 0.21
set label "Γ ≤ (Δ - 1)"     at graph 0.20, graph 0.76

plot "wb_topdown_polytope_many.csv" using 1:2 with points pt 5 ps 0.175 notitle, \
     (x+1)/x with lines axes x1y1 dt 1 lw 0.5  notitle, \
     (x-1)/1 with lines axes x1y1 dt 1 lw 0.5  notitle

