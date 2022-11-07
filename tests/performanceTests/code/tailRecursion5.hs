-- We need tail recursion to avoid stack overflow

main :: IO ();
main = do {
    tinySleep 2000;

    for 0 32 put0; -- 2^5
};

for :: Int -> Int -> (Int -> IO ()) -> IO ();
for i n fun = case i < n of {
    False -> putStr "";
    True -> do {
        fun i;
        for (i + 1) n fun;
    };
};

put0 :: Int -> IO ();
put0 n = putInt 0;
