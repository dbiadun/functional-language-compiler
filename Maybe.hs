data Maybe a = Just a | Nothing;

just :: Maybe Int -> Int;
just x = case x of {
    Just n -> n;
    Nothing -> 0;
};
