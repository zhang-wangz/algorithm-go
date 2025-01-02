# hashmap
class AuthenticationManager:

    def __init__(self, timeToLive: int):
        self.ttl = timeToLive
        self.map = {}

    def generate(self, tokenId: str, currentTime: int) -> None:
        self.map[tokenId] = currentTime + self.ttl

    def renew(self, tokenId: str, currentTime: int) -> None:
        if tokenId in self.map and currentTime < self.map[tokenId]:
            self.map[tokenId] = self.ttl + currentTime


    def countUnexpiredTokens(self, currentTime: int) -> int:
        cnt = 0
        for k, v in self.map.items():
            if v > currentTime:
                cnt += 1
        return cnt


# Your AuthenticationManager object will be instantiated and called as such:
# obj = AuthenticationManager(timeToLive)
# obj.generate(tokenId,currentTime)
# obj.renew(tokenId,currentTime)
# param_3 = obj.countUnexpiredTokens(currentTime)