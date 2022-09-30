data Maybe a = Just a | Nothing;

fromJust :: Maybe a -> a;
fromJust x = case x of {
    Just val -> val;
};
