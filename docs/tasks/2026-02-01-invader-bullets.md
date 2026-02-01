Time to build invader bullets. First, we will deal with their basic set-up. Shooting and collision will come later.

- InvaderBullets have same dimensions as PlayerBullets. They move down rather up. They have their own color setting.
- InvaderBullet should support Rectangler.
- Similarly to PlayerBullet, we will use a pool of size numInvaderBullets InvaderBullets. InvaderBullet has an `active` instance variable initially `false`.

Implement the struct with movement, drawing and rectangle methods, and implement the pool in Game. Implement unit tests for movement and pool creation.

---

The game functions without a loop initializing player bullets, because False is the zero value for boolean Active. So why do we need an initializing loop for invader bullets?

---

Now, we need invaders to fire bullets.

At intervals on invaderShootDelay, all invaders randomly decide shoot, each with probability invaderShootChance per cent.

If an invader decides to shoot, the game takes an inactive bullet from the pool and makes it active. (If no bullets are inactive, nothing further happens this frame; this is not an error.) As you can see, this part is analogous to player bullets.

When an invader shoots, the bullet appears at the bottom centre of that invader.

Note that invader bullets pass over invaders without interacting with them.

Plan the change and unit tests for it.
