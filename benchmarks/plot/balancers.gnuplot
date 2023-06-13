load "benchmarks/plot/colors.gnuplot"

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

set style line Median      dashtype 1 ps 2 lw 2 pt  11 pn 2 lc rgb COLOR_BLACK
set style line Height      dashtype 4 ps 2 lw 2 pt  10 pn 2 lc rgb COLOR_GREEN
set style line HalfSize    dashtype 1 ps 2 lw 2 pt   9 pn 2 lc rgb COLOR_RED
set style line LogSize     dashtype 5 ps 2 lw 2 pt   8 pn 2 lc rgb COLOR_YELLOW
set style line HalfWeight  dashtype 1 ps 2 lw 2 pt  13 pn 2 lc rgb COLOR_BLUE
set style line LogWeight   dashtype 5 ps 2 lw 2 pt  12 pn 2 lc rgb COLOR_CYAN
set style line Cost        dashtype 1 ps 2 lw 2 pt  14 pn 2 lc rgb COLOR_BROWN 
set style line DSW         dashtype 3 ps 2 lw 2 pt   6 pn 2 lc rgb COLOR_PINK

##################################################################
#
#           GROUPS
#
##################################################################

Partition = "Median Height HalfSize LogSize HalfWeight LogWeight Cost"
All       = "Median Height HalfSize LogSize HalfWeight LogWeight Cost DSW"
Limited   = "Median Height HalfWeight DSW"

GROUPS = "Partition All Limited"

DISTRIBUTIONS = "Uniform"


do for [GROUP in GROUPS] {

    ##################################################################
    #
    #           PARTITION COUNT
    #
    ##################################################################

    MEASUREMENT = "PartitionCount"

    set xlabel "{Size × 10^4}"
    set ylabel ""

    set title "{/:Bold Partition Count} / Size"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('PartitionCount')/column('Size')"

    set format y2 "%.2f"

    DATA = "benchmarks/data/balancers/measurements"

    SMOOTH = "sbezier"
    load "benchmarks/plot/balancers/lines.gnuplot"

    ##################################################################
    #
    #           TOTAL PARTITION DEPTH
    #
    ##################################################################

    MEASUREMENT = "TotalPartitionDepth"

    set xlabel "{Size × 10^4}"
    set ylabel ""

    set title "{/:Bold Partition Depth} / Size"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('PartitionDepth')/column('Size')"

    set format y2 "%.2f"

    DATA = "benchmarks/data/balancers/measurements"

    SMOOTH = "sbezier"
    load "benchmarks/plot/balancers/lines.gnuplot"


    ##################################################################
    #
    #           AVERAGE PARTITION DEPTH
    #
    ##################################################################

    MEASUREMENT = "AveragePartitionDepth"

    set xlabel "{Size × 10^4}"
    set ylabel ""

    set title "{/:Bold Partition Depth } / Partition Count"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('PartitionDepth')/column('PartitionCount')"

    set format y2 "%.2f"

    DATA = "benchmarks/data/balancers/measurements"

    SMOOTH = "sbezier"
    load "benchmarks/plot/balancers/lines.gnuplot"

    ##################################################################
    #
    #           ROTATIONS
    #
    ##################################################################

    MEASUREMENT = "Rotations"

    set xlabel "{Size × 10^4}"
    set ylabel ""

    set title "{/:Bold Rotations}"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('Rotations')"

    set format y2 "%.0f"

    DATA = "benchmarks/data/balancers/measurements"

    SMOOTH = "sbezier"
    load "benchmarks/plot/balancers/lines.gnuplot"


    ##################################################################
    #
    #           MAXIMUM PATH LENGTH
    #
    ##################################################################

    MEASUREMENT = "MaximumPathLength"

    set xlabel "Size × 10^4"
    set ylabel ""

    set title "{/:Bold Maximum Path Length } / log_2Size"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('MaximumPathLength')/log2(column('Size'))"

    set format y2 "%.2f"

    DATA = "benchmarks/data/balancers/measurements"

    SMOOTH = "sbezier"
    load "benchmarks/plot/balancers/lines.gnuplot"


    ##################################################################
    #
    #           AVERAGE PATH LENGTH
    #
    ##################################################################

    MEASUREMENT = "AveragePathLength"

    set xlabel "Size × 10^4"
    set ylabel ""

    set title "{/:Bold Average Path Length } / log_2Size"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('AveragePathLength')/log2(column('Size'))"

    set format y2 "%.2f"

    DATA = "benchmarks/data/balancers/measurements"

    SMOOTH = "sbezier"
    load "benchmarks/plot/balancers/lines.gnuplot"

    ##################################################################
    #
    #           DURATION
    #
    ##################################################################

    MEASUREMENT = "Duration"

    set xlabel "Size × 10^5"
    set ylabel ""

    set title "{/:Bold Duration } / Size"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('Duration')/column('Size')"

    set format y2 "%.0f"

    DATA = "benchmarks/data/balancers/benchmarks"

    SMOOTH = "sbezier"
    load "benchmarks/plot/balancers/lines.gnuplot"

    SMOOTH = "unique"
    load "benchmarks/plot/balancers/lines.gnuplot"
}
