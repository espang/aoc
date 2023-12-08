#!/usr/bin/env bb
(require
 '[clojure.math.numeric-tower :as math]
 '[clojure.string :as str])

(defn parse-line [l]
  (let [pattern #"(\w\w\w) = \((\w\w\w), (\w\w\w)\)"
        [_ start left right] (re-matches pattern l)]
    [start [left right]]))

(defn parse-network [network]
  (->> (str/split-lines network)
       (map parse-line)
       (into {})))

(defn parse-input [s]
  (let [[instructions network] (str/split s #"\n\n")]
    [instructions
     (parse-network network)]))

(defn move [network instruction position]
  (if-let [[left right] (get network position nil)]
    (case instruction
      \R right
      \L left)
    (throw (Exception. (str "unknown position: '" position "'")))))

(defn cycle-len [instructions network from]
  (loop [steps         0
         instructions' (cycle (seq instructions))
         position      from]
    (if (and (str/ends-with? position "Z")
             (zero? (mod steps (count instructions))))
      steps
      (recur (inc steps)
             (rest instructions')
             (move network (first instructions') position)))))

(defn solve-part2 [[instructions network]]
  (let [cycle-len' (partial cycle-len instructions network)]
    (->> (keys network)
         (filter #(str/ends-with? % "A"))
         (map cycle-len')
         (reduce (fn [acc v] (math/lcm acc v)) 1))))

(comment
  (def content (slurp "../../inputs/2023_8.txt"))
  (defn part1 [[instructions network]] (cycle-len instructions network "AAA"))
  (part1 (parse-input content))
  ;; part2
  (time
   (solve-part2 (parse-input content))))
