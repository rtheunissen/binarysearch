array AVLRedBlack[2]            = [ "AVLBottomUp", "RedBlackBottomUp" ]
array AVLWeak[4]                = [ "AVLBottomUp", "RedBlackBottomUp", "AVLWeakBottomUp", "AVLWeakTopDown" ]
array AVLRelaxed[4]             = [ "AVLRelaxedBottomUp", "AVLRelaxedTopDown", "RedBlackRelaxedBottomUp", "RedBlackRelaxedTopDown" ]
array RankBalanced[4]           = [ "AVLWeakBottomUp", "AVLWeakTopDown", "AVLRelaxedBottomUp", "RedBlackRelaxedTopDown" ]
array WeightBalanced[4]         = [ "LBSTBottomUp", "LBSTTopDown", "LBSTJoinBased", "LBSTRelaxed" ]
array Probabilistic[4]          = [ "TreapTopDown", "TreapFingerTree", "Randomized", "Zip" ]
array Other[2]                  = [ "AVLBottomUp", "Conc" ]
array Combination[4]            = [ "AVLRelaxedBottomUp", "AVLWeakBottomUp", "LBSTBottomUp", "TreapTopDown" ]
array CombinationSplay[5]       = [ "AVLRelaxedBottomUp", "AVLWeakBottomUp", "LBSTBottomUp", "TreapTopDown", "Splay" ]
array SizeOnly[4]               = [ "LBSTBottomUp", "LBSTRelaxed", "Randomized", "Splay" ]
array RedBlack[6]               = [ "AVLBottomUp", "AVLWeakBottomUp", "LBSTBottomUp", "RedBlackBottomUp", "RedBlackRelaxedBottomUp", "RedBlackRelaxedTopDown" ]

array Groups[2] = [ "AVLRedBlack",  "AVLWeak" ]

array Operations[7] = [ "Insert", "InsertPersistent", "InsertDelete", "InsertDeletePersistent", "InsertDeleteCycles", "InsertDeleteCyclesPersistent", "SplitJoin" ]
array Operations[5] = [ "Insert",  "InsertDelete", "InsertDeletePersistent", "InsertDeleteCycles", "InsertDeleteCyclesPersistent" ]

array Distributions[5] = [ "Uniform", "Normal", "Skewed", "Zipf", "Maximum" ]


do for [Operation=1:|Operations|] {

    OPERATION = Operations[Operation]

    csv = sprintf("benchmarks/csv/%s/", OPERATION)

    do for [Group=1:|Groups|] {

        GROUP = Groups[Group]

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

        smooth = "sbezier"
        load "benchmarks/plot/operations_lines.gnuplot"

        smooth = "unique"
        load "benchmarks/plot/operations_lines.gnuplot"

        smooth = "cumulative"
        load "benchmarks/plot/operations_lines.gnuplot"
    }
}