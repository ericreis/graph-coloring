# Graph COloring

## Benchmarks results

<table>
  <tr>
    <th rowspan="2">Instances<br></th>
    <th rowspan="2">|v|</th>
    <th rowspan="2">ρ</th>
    <th colspan="3">Backtrack (g.V)</th>
    <th colspan="3">Backtrack (k+1)</th>
    <th colspan="4">Branch and Bound</th>
    <th>Heuristic</th>
    <th></th>
    <th rowspan="2">Sol</th>
  </tr>
  <tr>
    <td>Execution time</td>
    <td>Explored solutions</td>
    <td>K</td>
    <td>Execution time</td>
    <td>Explored solutions</td>
    <td>K</td>
    <td>Execution time</td>
    <td>Explored solutions</td>
    <td>Branched solution</td>
    <td>K</td>
    <td>Execution time</td>
    <td>K</td>
  </tr>
  <tr>
    <td>test1</td>
    <td>3</td>
    <td>33.333%</td>
    <td>41.387µs</td>
    <td>6</td>
    <td>3</td>
    <td>40.38µs</td>
    <td>4</td>
    <td>3</td>
    <td>42.889µs</td>
    <td>0</td>
    <td>2</td>
    <td>3</td>
    <td>7.395µs</td>
    <td>3</td>
    <td>3</td>
  </tr>
  <tr>
    <td>test2</td>
    <td>4</td>
    <td>31.250%</td>
    <td>87.161µs</td>
    <td>36</td>
    <td>3</td>
    <td>61.906µs</td>
    <td>11</td>
    <td>3</td>
    <td>51.592µs</td>
    <td>0</td>
    <td>2</td>
    <td>3</td>
    <td>8.142µs</td>
    <td>3</td>
    <td>3</td>
  </tr>
  <tr>
    <td>test3</td>
    <td>5</td>
    <td>28.000%</td>
    <td>235.915µs</td>
    <td>80</td>
    <td>3</td>
    <td>73.871µs</td>
    <td>14</td>
    <td>3</td>
    <td>114.795µs</td>
    <td>0</td>
    <td>10</td>
    <td>3</td>
    <td>9.616µs</td>
    <td>3</td>
    <td>3</td>
  </tr>
  <tr>
    <td>myciel3</td>
    <td>11</td>
    <td>16.529%</td>
    <td>722.450ms</td>
    <td>249,460</td>
    <td>4</td>
    <td>13.187ms</td>
    <td>4,668</td>
    <td>4</td>
    <td>3.387ms</td>
    <td>80</td>
    <td>987</td>
    <td>4</td>
    <td>11.467µs</td>
    <td>4</td>
    <td>4</td>
  </tr>
  <tr>
    <td>myciel4</td>
    <td>23</td>
    <td>13.422%</td>
    <td>~72h</td>
    <td>48,694,100,000</td>
    <td>-</td>
    <td>44m52.565s</td>
    <td>787,863,524</td>
    <td>5</td>
    <td>15.257s</td>
    <td>505,200</td>
    <td>3,699,368</td>
    <td>5</td>
    <td>23.854µs</td>
    <td>5</td>
    <td>5</td>
  </tr>
  <tr>
    <td>myciel5</td>
    <td>47</td>
    <td>10.684%</td>
    <td>-</td>
    <td>-</td>
    <td>-</td>
    <td>-</td>
    <td>-</td>
    <td>-</td>
    <td> ~7h</td>
    <td>4,237,329,722</td>
    <td>341,000,000</td>
    <td>-</td>
    <td>49.131µs</td>
    <td>6</td>
    <td>6</td>
  </tr>
  <tr>
    <td>queen5_5</td>
    <td>25</td>
    <td>51.200%</td>
    <td>9.355s</td>
    <td>6,369</td>
    <td>5</td>
    <td>25.622ms</td>
    <td>1,692</td>
    <td>5</td>
    <td>564.316µs</td>
    <td>0</td>
    <td>9</td>
    <td>5</td>
    <td>33.288µs</td>
    <td>5</td>
    <td>5</td>
  </tr>
  <tr>
    <td>queen6_6</td>
    <td>36</td>
    <td>44.753%</td>
    <td></td>
    <td></td>
    <td>-</td>
    <td>3m42.257s</td>
    <td>439,764</td>
    <td>7</td>
    <td>5.141s</td>
    <td>3</td>
    <td>202,762</td>
    <td>7</td>
    <td>54.639µs</td>
    <td>9</td>
    <td>7</td>
  </tr>
  <tr>
    <td>queen7_7</td>
    <td>49</td>
    <td>39.650%</td>
    <td></td>
    <td></td>
    <td>-</td>
    <td>9m1.7289s</td>
    <td>3,617,667</td>
    <td>-</td>
    <td>12.487s</td>
    <td>59</td>
    <td>371.844</td>
    <td>7</td>
    <td>95.529µs</td>
    <td>12</td>
    <td>7</td>
  </tr>
</table>

## Results details

### test1.col

#### g.v

```
sol=[3 2 1] min=3
exploredSolutions=6
2018/11/24 18:00:19 backtrack took 41.387µs
```

#### k+1

```
sol=[2 3 1] min=3
exploredSolutions=4
2018/11/24 18:04:22 backtrack took 40.38µs
```

#### bab

```
sol=[1 2 3] min=3
exploredSolutions=0
branchedSolutions=2
2018/11/27 01:01:49 branch and bound took 42.889µs
```

#### greedy

```
sol=[1 3 2] k=3
2018/11/24 18:06:16 greedy took 7.395µs
```


### test2.col

#### g.v

```
sol=[4 3 4 1] min=3
exploredSolutions=36
2018/11/24 18:01:25 backtrack took 87.161µs
```

#### k+1

```
sol=[2 3 2 1] min=3
exploredSolutions=11
2018/11/24 18:03:23 backtrack took 61.906µs
```

#### bab

```
sol=[1 2 1 3] min=3
exploredSolutions=0
branchedSolutions=2
2018/11/27 01:01:32 branch and bound took 51.592µs
```

#### greedy

```
sol=[3 1 3 2] k=3
2018/11/24 18:06:46 greedy took 8.142µs
```


### test3.col

#### g.v

```
sol=[5 5 4 3 1] min=3
exploredSolutions=80
2018/11/24 18:02:29 backtrack took 235.915µs
```


#### k+1

```
sol=[2 2 3 4 1] min=3
exploredSolutions=14
2018/11/24 18:02:57 backtrack took 73.871µs
```

#### bab

```
sol=[1 1 2 3 2] min=3
exploredSolutions=0
branchedSolutions=10
2018/11/27 01:00:00 branch and bound took 114.795µs
```

#### greedy

```
sol=[1 1 2 3 2] k=3
2018/11/24 18:07:36 greedy took 9.616µs
```

### myciel3.col

#### g.V

```
sol=[11 10 11 10 9 11 10 11 10 9 1] min=4
exploredSolutions=249460
2018/11/27 01:13:08 backtrack took 722.450603ms
```

#### k+1

```
sol=[2 3 2 3 4 4 4 2 3 4 1] min=4
exploredSolutions=4668
2018/11/27 01:12:49 backtrack took 13.187563ms
```

#### bab

```
sol=[2 3 2 3 1 2 3 2 3 1 4] min=4
exploredSolutions=80
branchedSolutions=987
2018/11/27 00:59:34 branch and bound took 3.387112ms
```

#### greedy

```
sol=[1 2 3 2 1 3 2 3 2 4 1] k=4
2018/11/24 18:08:56 greedy took 11.467µs
```

### myciel4.col

#### g.V

```
exploredSolutions=48694100000
~72h
```

#### k+1

```
sol=[2 3 2 3 4 4 4 2 3 4 5 5 5 5 5 5 4 4 2 3 4 5 1] min=5
exploredSolutions=787863524
2018/11/16 23:16:11 backtrack took 44m52.565025305s
```

#### bab

```
sol=[2 3 2 3 4 4 4 2 3 4 1 2 3 2 3 4 4 4 2 3 4 1 5] min=5
exploredSolutions=505200
branchedSolutions=3699368
2018/11/27 00:59:16 branch and bound took 15.257577355s
```

#### greedy

```
sol=[1 2 3 2 1 3 4 3 3 4 1 5 2 3 2 4 3 2 3 2 4 2 1] k=5
2018/11/24 18:08:22 greedy took 23.854µs
```

### myciel5.col

#### g.V

```
-
```

#### k+1

```
exploredSolutions=534800000
~1h
```

#### bab

```
exploredSolutions=341000000
branchedSolutions=4237329722
~7h
```

#### greedy

```
sol=[2 1 3 1 2 2 5 3 3 2 1 4 4 3 3 4 6 5 3 3 5 4 1 2 5 3 3 2 2 5 3 3 2 4 2 4 3 3 2 2 4 3 3 2 4 2 1] k=6
2018/11/24 18:09:14 greedy took 49.131µs
```

### queen5_5

#### g.V

```
sol=[5 4 3 2 1 3 2 1 5 4 1 5 4 3 2 4 3 2 1 5 2 1 5 4 3] min=5
exploredSolutions=6369
2018/11/19 22:01:24 backtrack took 9.354607749s
```

#### k+1

```
sol=[2 3 4 5 1 5 1 2 3 4 3 4 5 1 2 1 2 3 4 5 4 5 1 2 3] min=5
exploredSolutions=1692
2018/11/19 22:44:58 backtrack took 25.622095ms
```

#### bab

```
sol=[1 2 3 4 5 3 4 5 1 2 5 1 2 3 4 2 3 4 5 1 4 5 1 2 3] min=5
exploredSolutions=0
branchedSolutions=9
2018/11/27 01:02:07 branch and bound took 564.316µs
```

#### greedy

```
sol=[2 3 4 1 5 1 5 2 3 4 3 4 1 5 2 5 2 3 4 1 4 1 5 2 3] k=5
2018/11/24 18:10:45 greedy took 33.288µs
```

### queen6_6

#### g.v

```

```

#### k+1

```
sol=[2 3 4 5 6 7 7 5 6 3 2 1 6 4 7 1 5 3 3 1 5 4 7 2 5 6 3 2 1 4 4 2 1 7 3 6] min=7
exploredSolutions=439764
2018/11/19 22:07:43 backtrack took 3m42.257375417s
```

#### bab

```
sol=[1 2 3 4 5 6 3 4 5 6 7 1 5 6 7 1 2 3 7 1 2 3 4 5 2 3 4 5 6 7 4 5 6 7 1 2] min=7
exploredSolutions=3
branchedSolutions=202762
2018/11/27 01:02:32 branch and bound took 5.1413681s
```

#### greedy

```
sol=[8 2 7 5 1 9 5 1 3 6 4 7 9 6 4 1 2 3 1 5 2 3 6 4 2 3 1 4 5 8 6 4 5 7 3 2] k=9
2018/11/24 18:11:25 greedy took 54.639µs
```

### queen7_7

#### g.v

```

```

#### k+1

```
sol=[2 3 4 5 6 7 1 7 1 2 3 4 5 6 5 6 7 1 2 3 4 3 4 5 6 7 1 2 1 2 3 4 5 6 7 6 7 1 2 3 4 5 4 5 6 7 1 2 3] min=7
exploredSolutions=3617667
backtrack took 9m1.728700528s
```

#### bab

```
sol=[1 2 3 4 5 6 7 3 4 5 6 7 1 2 5 6 7 1 2 3 4 7 1 2 3 4 5 6 2 3 4 5 6 7 1 4 5 6 7 1 2 3 6 7 1 2 3 4 5] min=7
exploredSolutions=59
branchedSolutions=371844
2018/11/27 01:03:03 branch and bound took 12.487382662s
```

#### greedy

```
sol=[12 4 3 8 6 5 7 8 5 6 7 1 3 4 9 1 2 3 4 6 5 7 3 4 1 5 2 9 4 6 5 2 3 1 8 11 2 1 6 7 4 10 10 7 8 4 2 9 6] k=12
2018/11/24 18:11:42 greedy took 95.529µs
```