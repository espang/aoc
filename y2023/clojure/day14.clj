(require '[clojure.string :as str])

(defn parse [s] (mapv vec (str/split-lines s)))
(defn shape [b] [(count (nth b 0)) (count b)])

(defn swap-until-stop [b [drow dcol] [row col]]
  (let [coord [row col]
        swapc [(+ row drow) (+ col dcol)]]
    (if (not= \O (get-in b coord))
      b
      (if (not= \. (get-in b swapc \#))
        b
        (recur (-> b
                   (assoc-in swapc \O)
                   (assoc-in coord \.))
               [drow dcol]
               swapc)))))

(defn bubble-column [b col reverse]
  (let [[_ rows] (shape b)
        change (if reverse [1 0] [-1 0])
        rng    (if reverse
                 (range (dec rows) -1 -1)
                 (range 0 rows))]
    (reduce (fn [b' row] (swap-until-stop b' change [row col])) b rng)))

(defn bubble-row [b row reverse]
  (let [[cols _] (shape b)
        change (if reverse [0 1] [0 -1])
        rng    (if reverse
                 (range (dec cols) -1 -1)
                 (range 0 cols))]
    (reduce (fn [b' col] (swap-until-stop b' change [row col])) b rng)))

(defn tilt [direction b]
  (let [[cols rows] (shape b)
        reducer (case direction
                  :north (fn [b' col] (bubble-column b' col false))
                  :south (fn [b' col] (bubble-column b' col true))
                  :west (fn [b' row] (bubble-row b' row false))
                  :east (fn [b' row] (bubble-row b' row true)))]
    (reduce reducer b (range 0 cols))))

(defn score [b]
  (let [[cols rows] (shape b)]
    (reduce +
            (for [x (range cols)
                  y (range rows)
                  :when (= \O (get-in b [y x]))]
              (- rows y)))))

(defn boards-seq [board]
  (let [board' (->> board
                    (tilt :north)
                    (tilt :west)
                    (tilt :south)
                    (tilt :east))]
    (cons board'
          (lazy-seq (boards-seq board')))))

(defn- find-cycle [coll]
  (let [value  (fn [idx] (nth coll idx))
        hare-index
        (loop [tortoise-index 0
               hare-index     1]

          (if (not= (value tortoise-index) (value hare-index))
            (recur (inc tortoise-index) (inc (inc hare-index)))
            hare-index))
        [mu tortoise-index']
        (loop [mu              0
               tortoise-index' 0
               hare-index'     (inc hare-index)]
          (if (not= (value tortoise-index') (value hare-index'))
            (recur (inc mu) (inc tortoise-index') (inc hare-index'))
            [mu tortoise-index']))
        lam
        (loop [lam 1
               hare-index'' (inc tortoise-index')]
          (if (not= (value tortoise-index') (value hare-index''))
            (recur (inc lam) (inc hare-index''))
            lam))]
    [mu lam]))

(defn solve-2
  " Solve the problem using Floyd's tortoise and hare algorithm
   https://en.wikipedia.org/wiki/Cycle_detection"
  [board]
  (let [scores (boards-seq board)
        [mu lam] (find-cycle scores)
        position (+ mu (mod (- 1000000000 (inc mu)) lam))]
    (println mu ":" lam)
    (nth scores position)))

(comment
  (def content (slurp "../../inputs/2023_14.txt"))
  ;; part 1
  (score (tilt :north (parse content)))
  ;; part 2
  (time
   (score (solve-2 (parse content)))))