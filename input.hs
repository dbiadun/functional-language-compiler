import "Maybe.hs";
import "List.hs";
-- Line comment

{- Multiline comment
-- Still a comment
-}

id :: b -> b;
id x = x;

append :: a -> List a -> List a;
append x = (id appendToList) x;

add :: Int -> Int -> Int;
add a b = a + b;

sum :: List Int -> Int;
sum = foldl add 0;

sumOfFirsts :: List (List Int) -> Int;
sumOfFirsts l = sum (map just (listOfFirsts l));

boolToInt :: Bool -> Int;
boolToInt b = case b of {
    False -> 0;
    True -> 1;
};

numbersFrom :: Int -> List Int;
numbersFrom n = Cons n (numbersFrom (n + 1));

numbers :: List Int;
numbers = numbersFrom 0;

sampleSumOfFirsts :: Int;
sampleSumOfFirsts = sumOfFirsts (Cons (Cons 4 (Cons 7 Nil)) (Cons Nil (Cons (Cons 2 (Cons 3 Nil)) Nil)));

sampleLazyOperators :: Int;
sampleLazyOperators = just (fst (Cons 3 (Cons (1/0) Nil)));

sampleTake :: Int; -- Checks lazy evaluation of infinite structures like `numbers`
sampleTake = sum (take 4 (Cons 3 (Cons 2 (Cons 4 (Cons 6 (Cons 7 Nil)))))); -- 15

main :: Int;
main = sampleTake;