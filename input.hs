data Maybe a = Just a | Nothing;
data List a = Cons a (List a) | Nil;

first :: List a -> Maybe a;
first x = case x of {
    Cons x xs -> Just x;
    Nil -> Nothing;
};

second :: List a -> Maybe a;
second x = case x of {
    Cons x xs -> first xs;
    Nil -> Nothing;
};

firstTwo :: List a -> Maybe (List a);
firstTwo l = case (first l) of {
    Nothing -> Nothing;
    Just f -> case (second l) of {
        Nothing -> Nothing;
        Just s -> Just (Cons f (Cons s Nil));
    };
};

appendToList :: a -> (List a -> List a);
appendToList x xs = Cons x xs;

id :: b -> b;
id x = x;

append :: a -> List a -> List a;
append x = (id appendToList) x;

map :: (a -> b) -> List a -> List b;
map f l = case l of {
    Nil -> Nil;
    Cons x xs -> Cons (f x) (map f xs);
};

foldl :: (a -> b -> a) -> a -> List b -> a;
foldl fun acc l = case l of {
    Nil -> acc;
    Cons x xs -> (foldl fun) (fun acc x) xs;
};

listOfFirsts :: List (List a) -> List (Maybe a);
listOfFirsts = map first;

add :: Int -> Int -> Int;
add a b = a + b;

sum :: List Int -> Int;
sum = foldl add 0;

