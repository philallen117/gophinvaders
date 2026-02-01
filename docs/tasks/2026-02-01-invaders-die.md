The player's `score` is initially zero. The score text displays as described in config.

The text is for example "Score: 170".

Add the config above, add the score initialization and display to the game.

--- Already done, but problems with text display - see (./2026-02-01-text-handling.md) which had to be done first.

Now, we will support the mechanic of player bullets killing invaders and adding to the score.

When a player bullet collides with an invader:

- The invader is removed.
- The bullet returns to the pool and becomes inactive.
- The score increases by KILL_SCORE.

Plan the implementation. There will be other collisions to handle in game logic later on, so I prefer a design where collision logic is easy to reuse.

---

Carry out this plan. Include testing: in each frame,

- Each bullet kills at most one invader
- More than one bullet may strike (different) invaders.

---

At this point I realized it did not lint before, so I have ended up with a library version change in the same commit. Grrrr.

---

CheckCollision takes two Rectanglers. The problem with this is that  HandleBulletInvaderCollisions finds the rectangle of the bullet repeatedly inside the inner invaders loop. This is not necessary.
Change CheckCollision to take two four-tuples, and find the rectangle of the bullet in the outer loop of HandleBulletInvaderCollisions.
