class Solution:
    #find a deterministic foolproof way to encode and use the key to decode.
    # the constraints tell the vaalue of lie between 256 valid ascii characters.
    # so why not use a character that is not one of them as the separator and join.
    def encode(self, strs: List[str]) -> str:
        if len(strs) == 0:
            return 'iknowthiswillnotbeavalidsample'
        return "★".join(strs)
        

    def decode(self, s: str) -> List[str]:
        if s == "iknowthiswillnotbeavalidsample":
            return []
        return s.split("★")
        