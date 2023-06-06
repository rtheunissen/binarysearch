set terminal svg size 800,500 round dynamic

set samples 1000

set border 6 front lc '#000000' lt 1 lw 0

DEFAULT_FONT = ",16"
HEADER_FONT = ",24"

set style line 1000 dashtype 1 lw 1 pt 1 ps 0 lc "#EEEEEE"

set grid y2tics back linestyle 1000
set grid xtics back linestyle 1000

set title font HEADER_FONT
set title offset 0,6

set xlabel  font DEFAULT_FONT offset 0,-1
set ylabel  font DEFAULT_FONT offset -1,0

set key width 0
set key height 0
set key spacing 1.5
set key samplen 4
set key reverse Left
set key horizontal outside center top
set key font DEFAULT_FONT
set key at graph 0.5, screen 0.86

set autoscale xfix

unset xtics
unset ytics

unset mxtics
unset mytics

unset my2tics
unset mx2tics

unset x2tics
unset y2tics

set xtics font DEFAULT_FONT offset 0,-0.5
set xrange [0.8:*]

unset y2tics
set y2tics autofreq font DEFAULT_FONT offset 1,0
set y2range [*:*]

set datafile missing "NaN"


