# # date > myciel5_bab.log
# # go run main.go myciel5 bab >> myciel5_bab.log
# date > queen7_7_backtrack.log
# go run main.go queen7_7 backtrack >> queen7_7_backtrack.log
# date > myciel5_backtrack.log
# go run main.go myciel5 backtrack >> myciel5_backtrack.log


# for algo in heuristic+ heuristic meta
for algo in meta
do
  mkdir -p "logs/$algo"
  for instance in test1 test2 test3 myciel3 myciel4 myciel5 queen5_5 queen6_6 queen7_7 2-Insertions_3 homer fpsol2.i.1 fpsol2.i.2 fpsol2.i.3 inithx.i.1 qg.order100 qg.order30
  do
    go run main.go "$instance" $algo > "logs/$algo/$instance.log"
  done
done