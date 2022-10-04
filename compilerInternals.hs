import "List.hs";
import "Either.hs";

compilerListIntEmpty :: List Int;
compilerListIntEmpty = Nil;

compilerListAppend :: List a -> a -> List a;
compilerListAppend l elem = Cons elem l;

compilerEitherLeftString :: String -> Either String (List Int);
compilerEitherLeftString s = Left s;

compilerEitherRightList :: List Int -> Either String (List Int);
compilerEitherRightList l = Right l;

compilerErrorMaybeNil :: Maybe String;
compilerErrorMaybeNil = Nothing;

compilerErrorMaybeString :: String -> Maybe String;
compilerErrorMaybeString s = Just s;
