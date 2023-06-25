load "benchmarks/plot/colors.gnuplot"

##################################################################
#
#           STRATEGIES
#
##################################################################

AVLBottomUp                = 1
AVLJoinBased               = 2
AVLWeakTopDown             = 3
AVLWeakBottomUp            = 4
AVLWeakJoinBased           = 5
AVLRelaxedTopDown          = 6
AVLRelaxedBottomUp         = 7
RedBlackRelaxedBottomUp    = 8
RedBlackRelaxedTopDown     = 9
LBSTBottomUp               = 10
LBSTTopDown                = 11
LBSTJoinBased              = 12
LBSTRelaxed                = 13
TreapTopDown               = 14
TreapJoinBased             = 15
TreapFingerTree            = 16
Randomized                 = 17
Zip                        = 18
Splay                      = 19
Conc                       = 20
RedBlackBottomUp           = 21


set style line AVLBottomUp                  dashtype 1 ps 2 lw 2 pt   1 pn 2 lc rgb COLOR_BLACK
set style line AVLJoinBased                 dashtype 3 ps 2 lw 2 pt   2 pn 2 lc rgb COLOR_CYAN
set style line AVLWeakBottomUp              dashtype 1 ps 2 lw 2 pt   4 pn 2 lc rgb COLOR_BLUE
set style line AVLWeakTopDown               dashtype 5 ps 2 lw 2 pt   3 pn 2 lc rgb COLOR_GREEN
set style line AVLWeakJoinBased             dashtype 3 ps 2 lw 2 pt   5 pn 2 lc rgb COLOR_CYAN
set style line AVLRelaxedTopDown            dashtype 5 ps 2 lw 2 pt   6 pn 2 lc rgb COLOR_BLUE
set style line AVLRelaxedBottomUp           dashtype 1 ps 2 lw 2 pt   7 pn 2 lc rgb COLOR_YELLOW
set style line RedBlackBottomUp             dashtype 1 ps 2 lw 2 pt  12 pn 2 lc rgb COLOR_RED
set style line RedBlackRelaxedBottomUp      dashtype 1 ps 2 lw 2 pt   8 pn 2 lc rgb COLOR_GREEN
set style line RedBlackRelaxedTopDown       dashtype 5 ps 2 lw 2 pt   9 pn 2 lc rgb COLOR_CYAN
set style line LBSTBottomUp                 dashtype 1 ps 2 lw 2 pt  10 pn 2 lc rgb COLOR_GREEN
set style line LBSTTopDown                  dashtype 5 ps 2 lw 2 pt  11 pn 2 lc rgb COLOR_BLUE
set style line LBSTJoinBased                dashtype 3 ps 2 lw 2 pt  12 pn 2 lc rgb COLOR_YELLOW
set style line LBSTRelaxed                  dashtype 1 ps 2 lw 2 pt  13 pn 2 lc rgb COLOR_RED
set style line TreapTopDown                 dashtype 5 ps 2 lw 2 pt  14 pn 2 lc rgb COLOR_RED
set style line TreapJoinBased               dashtype 3 ps 2 lw 2 pt  15 pn 2 lc rgb COLOR_CYAN
set style line TreapFingerTree              dashtype 1 ps 2 lw 2 pt  16 pn 2 lc rgb COLOR_BLUE
set style line Randomized                   dashtype 1 ps 2 lw 2 pt  17 pn 2 lc rgb COLOR_YELLOW
set style line Zip                          dashtype 1 ps 2 lw 2 pt  18 pn 2 lc rgb COLOR_GREEN
set style line Splay                        dashtype 1 ps 2 lw 2 pt  19 pn 2 lc rgb COLOR_PINK
set style line Conc                         dashtype 3 ps 2 lw 2 pt  20 pn 2 lc rgb COLOR_PURPLE


##################################################################
#
#           GROUPS
#
##################################################################

AVLRedBlack            = "AVLBottomUp RedBlackBottomUp"
AVLWeakRedBlack        = "AVLBottomUp RedBlackBottomUp AVLWeakBottomUp AVLWeakTopDown"
AVLWeak                = "AVLBottomUp AVLWeakJoinBased AVLWeakBottomUp AVLWeakTopDown"
AVLConc                = "AVLBottomUp AVLWeakBottomUp AVLWeakTopDown AVLRelaxedTopDown AVLRelaxedBottomUp Conc"
AVLRelaxed             = "AVLRelaxedBottomUp AVLRelaxedTopDown RedBlackRelaxedBottomUp RedBlackRelaxedTopDown"

RedBlack               = "AVLBottomUp AVLWeakBottomUp RedBlackBottomUp RedBlackRelaxedBottomUp RedBlackRelaxedTopDown"
RankBalanced           = "AVLWeakBottomUp AVLWeakTopDown AVLRelaxedBottomUp RedBlackRelaxedTopDown"
HeightBalanced         = "AVLBottomUp RedBlackBottomUp AVLWeakBottomUp AVLRelaxedBottomUp Conc"

WeightBalanced         = "LBSTBottomUp LBSTTopDown LBSTJoinBased LBSTRelaxed"
Probabilistic          = "TreapTopDown TreapFingerTree Randomized Zip"
Combination            = "AVLRelaxedBottomUp AVLWeakBottomUp LBSTBottomUp TreapTopDown"
CombinationSplay       = "AVLRelaxedBottomUp AVLWeakBottomUp LBSTBottomUp TreapTopDown Splay"
SizeOnly               = "LBSTBottomUp LBSTRelaxed Randomized Splay"

GROUPS = "AVLWeakRedBlack AVLRelaxed RedBlack"

OPERATIONS = "InsertPersistent InsertDeleteCyclesPersistent Insert InsertDeleteCycles InsertDelete InsertDeletePersistent InsertDeleteCycles"
OPERATIONS = "Insert InsertDeleteCycles InsertDelete"

DISTRIBUTIONS = "Uniform Normal Skewed Zipf Maximum"

do for [GROUP in GROUPS] {

    do for [OPERATION in OPERATIONS] {

        ##################################################################
        #
        #           Allocations
        #
        ##################################################################

        MEASUREMENT = 'Allocations'

        set xlabel "Operations × 10^5"
        set ylabel "{/:Bold Allocations } / log_2Size"

        DATA = sprintf("benchmarks/data/operations/measurements/%s", OPERATION)

        x = "column('Position')/(column('Scale')/10)"
        y = "column('Allocations')/column('Step')/log2(column('Size'))"

        set format y2 "%.2f"

        SMOOTH = "sbezier"
        load "benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "benchmarks/plot/operations/lines.gnuplot"

        ##################################################################
        #
        #           MAXIMUM PATH LENGTH
        #
        ##################################################################

        MEASUREMENT = 'MaximumPathLength'

        set xlabel "Operations × 10^5"
        set ylabel "{/:Bold Maximum Path Length } / log_2Size"

        x = "column('Position')/(column('Scale')/10)"
        y = "column('MaximumPathLength')/log2(column('Size'))"

        DATA = sprintf("benchmarks/data/operations/measurements/%s", OPERATION)

        set format y2 "%.2f"

        SMOOTH = "sbezier"
        load "benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "benchmarks/plot/operations/lines.gnuplot"

        ##################################################################
        #
        #           AVERAGE PATH LENGTH
        #
        ##################################################################

        MEASUREMENT = 'AveragePathLength'

        set xlabel "Operations × 10^5"
        set ylabel "{/:Bold AveragePathLength } / log_2Size"

        DATA = sprintf("benchmarks/data/operations/measurements/%s", OPERATION)

        set format y2 "%.2f"

        x = "column('Position')/(column('Scale')/10)"
        y = "column('AveragePathLength')/log2(column('Size'))"

        SMOOTH = "sbezier"
        load "benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "benchmarks/plot/operations/lines.gnuplot"

        ##################################################################
        #
        #           ROTATIONS
        #
        ##################################################################

        MEASUREMENT = "Rotations"

        set xlabel "{Operations × 10^5}"
        set ylabel "{/:Bold Rotations} / Operation"

        x = "column('Position')/(column('Scale')/10)"
        y = "column('Rotations')/column('Step')"

        DATA = sprintf("benchmarks/data/operations/measurements/%s", OPERATION)

        set format y2 "%.2f"

        SMOOTH = "sbezier"
        load "benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "benchmarks/plot/operations/lines.gnuplot"

        ##################################################################
        #
        #           DURATION
        #
        ##################################################################

        MEASUREMENT = 'Duration'

        set xlabel "Operations / 10^5"
        set ylabel "{/:Bold Duration } in milliseconds"

        x = "column('Position')/(column('Scale')/10)"
        y = "column('Duration')/1000000" # nano / 10^6 is 1k

        set format y2 "%.2f"

        DATA = sprintf("benchmarks/data/operations/benchmarks/%s", OPERATION)

        SMOOTH = "sbezier"
        load "benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "cumulative"
        load "benchmarks/plot/operations/lines.gnuplot"
    }
}