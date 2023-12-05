#!/usr/bin/env bb
(require '[clojure.string :as str])

(def content (slurp "../../inputs/2023_5.txt"))

(defn parse-int-list [s]
  (->> (str/split s #" ")
       (map #(Long/parseLong %))))

(defn parse-seeds [s]
  (-> (str/replace-first s "seeds: " "")
      parse-int-list))

(defn parse-map [s]
  (->> (str/split-lines s)
       (drop 1)
       (map parse-int-list)))

(defn transform [start [destination source range']]
  (when (<= source start (+ source (dec range')))
    (+ destination (- start source))))

(defn apply-map [map' seed]
  (reduce (fn [acc v]
            (if-some [new-value (transform acc v)]
              (reduced new-value)
              acc))
          seed
          map'))

(defn apply-all-maps [maps seed]
  [seed (reduce (fn [acc m] (apply-map m acc)) seed maps)])

(defn part1 [content]
  (let [lines (str/split content #"\n\n")
        seeds (parse-seeds (first lines))
        maps  (->> (rest lines)
                   (map parse-map))]
    (->> seeds
         (map (partial apply-all-maps maps)))))

(defn convert-seeds [seeds]
  (->> (partition 2 seeds)
       (map (fn [[start delta]] [start (+ start (dec delta))]))
       (sort-by first)))

(defn valid-seed [seed seed-ranges]
  (if-not (seq seed-ranges)
    false
    (let [[start end] (first seed-ranges)]
      (if (<= start seed end)
        true
        (recur seed (rest seed-ranges))))))

(defn convert-map [map']
  (->> map'
       (map (fn [[dest src range']]
              [dest
               (+ dest (dec range')) (- src dest)]))
       (sort-by first)))

(defn traverse [v map']
  (if-not (seq map')
    v
    (let [[start end delta] (first map')]
      (if (<= start v end)
        (+ v delta)
        (recur v (rest map'))))))

(defn traverse-all [v maps]
  (reduce (fn [v' map'] (traverse v' map')) v maps))

(defn part2 [content]
  (let [lines (str/split content #"\n\n")
        seeds (-> (first lines)
                  parse-seeds
                  convert-seeds)
        maps  (->> (rest lines)
                   (map parse-map)
                   (map convert-map)
                   reverse)
        [_ end _] (first (first maps))]
    (loop [i 0]
      (when (zero? (mod i 100000))
        (println "Handled " i))
      (if (> i end)
        -1
        (let [potential-seed (traverse-all i maps)]
          (if (valid-seed potential-seed seeds)
            i
            (recur (inc i))))))))

(comment
  ;; (time (apply min-key second (part1 test-input)))
  (time (apply min-key second (part1 content)))
  ;; (time (part2 test-input))
  (time (part2 content)))
