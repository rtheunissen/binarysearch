load "benchmarks/plot/colors.gnuplot"

array Partition[7] = [ \
    "Median", \
    "Height", \
    "HalfSize", \
    "LogSize", \
    "HalfWeight", \
    "LogWeight", \
    "Cost", \
]

array All[8] = [ \
    "Median", \
    "Height", \
    "HalfSize", \
    "LogSize", \
    "HalfWeight", \
    "LogWeight", \
    "Cost", \
    "DSW", \
]

array Limited[3] = [ \
    "Median", \
    "Height", \
    "DSW", \
]

array Groups[3] = [ "Partition", "All", "Limited" ]

array Distributions[1] = [ "Uniform" ]


##################################################################
#
#           BALANCERS
#
##################################################################

Median      = 1
Height      = 2
HalfSize    = 3
HalfWeight  = 4
LogSize     = 5
LogWeight   = 6
Cost        = 7
DSW         = 8


set style line Median      dashtype 1  ps 2 lw 2 pt  11 pn 2 lc rgb COLOR_BLACK
set style line Height      dashtype 4  ps 2 lw 2 pt  10 pn 2 lc rgb COLOR_BROWN

set style line HalfSize    dashtype 1  ps 2 lw 2 pt   9 pn 2 lc rgb COLOR_RED
set style line LogSize     dashtype 1  ps 2 lw 2 pt   8 pn 2 lc rgb COLOR_YELLOW

set style line HalfWeight  dashtype 1  ps 2 lw 2 pt  13 pn 2 lc rgb COLOR_BLUE
set style line LogWeight   dashtype 1  ps 2 lw 2 pt  12 pn 2 lc rgb COLOR_CYAN

set style line Cost        dashtype 1  ps 2 lw 2 pt  14 pn 2 lc rgb COLOR_GREEN
set style line DSW         dashtype 1  ps 2 lw 2 pt   6 pn 2 lc rgb COLOR_PINK

