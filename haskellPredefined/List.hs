import "Maybe.hs";

data List a = Cons a (List a) | Nil;

appendToList :: a -> (List a -> List a);
appendToList x xs = Cons x xs;

fst :: List a -> Maybe a;
fst x = case x of {
    Cons x xs -> Just x;
    Nil -> Nothing;
};

snd :: List a -> Maybe a;
snd x = case x of {
    Cons x xs -> fst xs;
    Nil -> Nothing;
};

last :: List a -> Maybe a;
last x = case x of {
    Cons x xs -> case xs of {
        Nil -> Just x;
        l -> last l;
    };
    Nil -> Nothing;
};

firstTwo :: List a -> Maybe (List a);
firstTwo l = case (fst l) of {
    Nothing -> Nothing;
    Just f -> case (snd l) of {
        Nothing -> Nothing;
        Just s -> Just (Cons f (Cons s Nil));
    };
};

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
listOfFirsts = map fst;

listOfLasts :: List (List a) -> List (Maybe a);
listOfLasts = map last;

take :: Int -> List a -> List a;
take n l = case n > 0 of {
    False -> Nil;
    True -> case l of {
        Nil -> Nil;
        Cons x xs -> Cons x (take (n - 1) xs);
    };
};
