(require '[clojure.string :as str])

(defn char->int [c] (- (int c) (int \0)))

(defn parse [s]
  (->> (str/split-lines s)
       (map #(mapv char->int %))
       (mapv vec)))

(defn on-board? [board [row col]]
  (let [rows (count board)
        cols (count (nth board 0))]
    (and (<= 0 row (dec rows))
         (<= 0 col (dec cols)))))

(defn move [[row col] direction]
  (case direction
    :north [(dec row) col]
    :south [(inc row) col]
    :west  [row (dec col)]
    :east  [row (inc col)]))

(defn next-positions [board [previous-directions position]]
  (let [three-equal-moves (and (= 3 (count previous-directions))
                               (apply = previous-directions))]
    (->> (case (first previous-directions)
           nil    [:north :south :west :east]
           :north (if three-equal-moves [:west :east] [:north :west :east])
           :south (if three-equal-moves [:west :east] [:south :west :east])
           :west  (if three-equal-moves [:north :south] [:north :south :west])
           :east  (if three-equal-moves [:north :south] [:north :south :east]))
         (map (fn [dir] [dir (move position dir)]))
         (filter (fn [[_ pos]] (on-board? board pos))))))

(defn make-last-directions [dir dirs]
  (if (<= (count dirs) 2)
    (conj dirs dir)
    (conj (take 2 dirs) dir)))

(defn walk [board]
  (let [start  [nil [0 0]]
        rows   (count board)
        cols   (count (nth board 0))
        target [(dec rows) (dec cols)]]
    (loop [queue         (conj (clojure.lang.PersistentQueue/EMPTY) [start 0])
           heatlosses    {}
           best-solution Integer/MAX_VALUE]
      (if (empty? queue)
        heatlosses
        (let [[current heatloss] (peek queue)
              [dirs position]       current
              best-solution      (if (= target position)
                                   (min heatloss best-solution)
                                   best-solution)]
          (if (and (< heatloss (get heatlosses current Integer/MAX_VALUE))
                   (<= heatloss best-solution))
            ;; found a better way --> continue
            (let [next-positions (next-positions board current)
                  next-elements  (map (fn [[dir' pos']]
                                        [[(make-last-directions dir' dirs)
                                          pos']
                                         (+ heatloss (get-in board pos'))])
                                      next-positions)]
              (recur (apply conj (pop queue)
                            next-elements)
                     (assoc heatlosses current heatloss)
                     best-solution))
            ;; found better solution before; skip
            (recur (pop queue) heatlosses best-solution)))))))

(defn next-positions-2 [board [[previous-direction n] position]]
  (->> (if (nil? previous-direction)
         [:south :east]
         (case previous-direction
           :north (if (< n 4)
                    [:north]
                    (if (= n 10) [:west :east] [:north :west :east]))
           :south (if (< n 4)
                    [:south]
                    (if (= n 10) [:west :east] [:south :west :east]))
           :west  (if (< n 4)
                    [:west]
                    (if (= n 10) [:north :south] [:north :south :west]))
           :east  (if (< n 4)
                    [:east]
                    (if (= n 10) [:north :south] [:north :south :east]))))
       (map (fn [dir] [dir (move position dir)]))
       (filter (fn [[_ pos]] (on-board? board pos)))))

(defn walk-2 [board]
  (let [start  [nil [0 0]]
        rows   (count board)
        cols   (count (nth board 0))
        target [(dec rows) (dec cols)]]
    (loop [queue         (conj (clojure.lang.PersistentQueue/EMPTY) [start 0])
           heatlosses    {}
           best-solution Integer/MAX_VALUE]
      (if (empty? queue)
        heatlosses
        (let [[current heatloss] (peek queue)
              [[dir n] position]       current
              best-solution      (if (= target position)
                                   (min heatloss best-solution)
                                   best-solution)]
          (if (and (< heatloss (get heatlosses current Integer/MAX_VALUE))
                   (<= heatloss best-solution))
            ;; found a better way --> continue
            (let [next-positions (next-positions-2 board current)
                  next-elements  (map (fn [[dir' pos']]
                                        [[(if (= dir dir')
                                            [dir  (inc n)]
                                            [dir' 1])
                                          pos']
                                         (+ heatloss (get-in board pos'))])
                                      next-positions)]
              (recur (apply conj (pop queue)
                            next-elements)
                     (assoc heatlosses current heatloss)
                     best-solution))
            ;; found better solution before; skip
            (recur (pop queue) heatlosses best-solution)))))))

(comment
  (def content (slurp "../../inputs/2023_17.txt"))
  (def board (parse content))

  ;; part 1 || ~ 400s
  (time (def losses (walk board)))
  (->> (filter (fn [[[_ [row col]] _]]
                 (and (= row 140) (= col 140)))
               losses)
       (apply min-key (fn [[_ v]] v)))

  ;; part 2 || ~ 160s
  (time (def losses-2 (walk-2 (parse content))))
  (->> (filter (fn [[[_ [row col]] _]]
                 (and (= row 140) (= col 140)))
               losses-2)
       (apply min-key (fn [[_ v]] v))))

