import "List.hs";
import "Maybe.hs";

main :: IO ();
main = do {
    sum <- return (sumOfLasts 21);
    putInt sum;
};

sumOfLasts :: Int -> Int; -- Checks the performance when many nodes are needed (garbage collection)
sumOfLasts n = sum (map just (listOfLasts (map nNumbers (nNumbers (n + 1)))));

just :: Maybe Int -> Int;
just x = case x of {
    Just n -> n;
    Nothing -> 0;
};

sum :: List Int -> Int;
sum = foldl add 0;

add :: Int -> Int -> Int;
add a b = a + b;

nNumbers :: Int -> List Int;
nNumbers n = take n numbers;

numbers :: List Int;
numbers = numbersFrom 0;

numbersFrom :: Int -> List Int;
numbersFrom n = Cons n (numbersFrom (n + 1));

listOfLasts :: List (List a) -> List (Maybe a);
listOfLasts = map last;
