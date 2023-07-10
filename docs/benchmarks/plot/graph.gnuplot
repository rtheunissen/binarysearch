set terminal svg size 400,250 dynamic round font "KaTeX_Main,10"

set samples 1000

set border 6 front lc '#000000' lt -1 lw 0

set title font "KaTeX_Main,10"

set style line 1000 dashtype 3 lw 0.2 ps 0 lc rgb "#333333"

#set grid y2tics back linestyle 1000
set grid xtics back linestyle 1000

set xlabel offset 0,0.5
set ylabel offset 0,0

set key width 0
set key height 0
set key spacing 1.25
set key samplen 4
set key reverse Left
set key horizontal outside center top
set key at graph 0.5, screen 1
set key font "KaTeX_Main,8"

unset xtics
unset ytics

unset mxtics
unset mytics
unset my2tics
unset mx2tics

unset x2tics
unset y2tics

set xtics offset 0,0.5
set xrange [0.8:*]

unset y2tics
set y2tics autofreq offset 0,0
set y2range [*:*]

set datafile missing "NaN"


