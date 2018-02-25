#!/bin/bash

gnuplot -e "initialstate='Random'" analysis/perInitialStatenlogn.gp
gnuplot -e "initialstate='Nearly Sorted'" analysis/perInitialStatenlogn.gp
gnuplot -e "initialstate='Reversed'" analysis/perInitialStatenlogn.gp
gnuplot -e "initialstate='Few Unique'" analysis/perInitialStatenlogn.gp

gnuplot -e "initialstate='Random'" analysis/perInitialState.gp
gnuplot -e "initialstate='Nearly Sorted'" analysis/perInitialState.gp
gnuplot -e "initialstate='Reversed'" analysis/perInitialState.gp
gnuplot -e "initialstate='Few Unique'" analysis/perInitialState.gp

gnuplot -e "algrtm='Bubble Sort'" analysis/perAlgorithm.gp
gnuplot -e "algrtm='Cocktail Sort'" analysis/perAlgorithm.gp
gnuplot -e "algrtm='Heap Sort'" analysis/perAlgorithm.gp
gnuplot -e "algrtm='Insertion Sort'" analysis/perAlgorithm.gp
gnuplot -e "algrtm='Native Sort'" analysis/perAlgorithm.gp
gnuplot -e "algrtm='Quick Sort'" analysis/perAlgorithm.gp
gnuplot -e "algrtm='Selection Sort'" analysis/perAlgorithm.gp
gnuplot -e "algrtm='Shell Sort'" analysis/perAlgorithm.gp
