set terminal svg size 400,200 dynamic round \
    font "docs/katex/fonts/KaTeX_Main-Regular.ttf,8"

set samples 1000

set border 6 front lc '#000000' lt 1 lw 0

set style line 1000 dashtype 1 lw 0.5 pt 1 ps 0 lc "#AAAAAAAA"

#set grid y2tics back linestyle 1000
#set grid xtics back linestyle 1000

set title offset 0,0

set xlabel offset 0,1
set ylabel offset 0,0

set key width 0
set key height 0
set key spacing 1.25
set key samplen 4
set key reverse Left
set key horizontal outside center top
set key at graph 0.5, screen 1

#set autoscale xfix

unset xtics
unset ytics

unset mxtics
unset mytics

unset my2tics
unset mx2tics

unset x2tics
unset y2tics

set xtics offset 0,0.5
set xrange [0.1:*]

unset y2tics
set y2tics autofreq offset 0,0
set y2range [*:*]

set datafile missing "NaN"


