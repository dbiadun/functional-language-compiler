
----------------------------------- PURE FUNCTIONS --------------------------------------

{- Bitwise left shift operation
-- 1. param - shifted number
-- 2. param - number of bits to shift
-}
shiftL :: Int -> Int -> Int;

{- Bitwise left shift operation
-- 1. param - shifted number
-- 2. param - number of bits to shift
-}
shiftR :: Int -> Int -> Int;

------------------------------------- INPUT/OUTPUT --------------------------------------

-- Changes a value into an IO action returning this value
return :: a -> IO a;

-- Prints a string (on microcontrollers it prints to the serial output)
putStr :: String -> IO ();

-- Prints an integer
putInt :: Int -> IO ();

---------------------------------------- STATE ------------------------------------------
-- State is organized as a mapping from strings to integers

{- Sets a value of a variable in the state
-- 1. param - variable name
-- 2. param - value to set
-}
stateSetInt :: String -> Int -> IO ();

{- Gets a variable from the state
-- 1. param - variable name
-- returns a value stored in the variable
-}
stateGetInt :: String -> IO Int;

-- Go functions
tinySleep :: Int -> IO ();
