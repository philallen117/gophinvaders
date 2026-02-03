I see HandlePlayerBulletShieldCollisions and HandleInvaderBulletShieldCollisions have a lot of code in common. Can these methods be generalized to reduce duplication?

---

For uniformity extra player shooting logic from Game.Update into a new method.
