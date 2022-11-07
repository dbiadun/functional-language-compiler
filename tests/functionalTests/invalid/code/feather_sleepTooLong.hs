main :: IO ();
main = do {
    tinySleep 2000;

    putStr "Start\n";
    tinySleep 5000;
    putStr "The end\n"; -- The output shouldn't match because we have a 5s timeout on waiting for microcontroller's response
};