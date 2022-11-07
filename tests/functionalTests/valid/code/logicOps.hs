main :: IO ();
main = do {
    b <- return (True && False || (True || False));
    case b of {
        True -> putStr "True";
        False -> putStr "False";
    };
};

