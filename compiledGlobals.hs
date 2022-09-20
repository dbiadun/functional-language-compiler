return :: a -> IO a;
putStr :: String -> IO ();
putInt :: Int -> IO ();
getLine :: IO String;

-- Go functions
tinySleep :: Int -> IO ();

-- TinyGo microcontroller functions
tinyLED :: IO Int;

tinyConfigure :: Int -> Int -> IO ();
tinyPinOutput :: IO Int;

tinyLow :: Int -> IO ();
tinyHigh :: Int -> IO ();
