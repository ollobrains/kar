package system

import (
    "math"

    "kar"
    "kar/arc"
    "kar/engine/mathutil"

    "github.com/mlange-42/arche/ecs"
)

// Collect handles items that can be picked up, e.g., if the player's bounding box overlaps an item.
//
// Example usage:
//   var collectSystem Collect
//   collectSystem.Init()
//   // in main loop:
//   collectSystem.Update()
//   collectSystem.Draw()
type Collect struct {
    // itemGravity controls how quickly items fall on the Y-axis
    itemGravity float64

    // sinSpace can be used to animate items in some sinusoidal pattern
    sinSpace []float64
    // precomputed length for quick modulus usage
    sinSpaceLen int

    // a temporary slice for items scheduled to be removed this frame
    itemsToRemove []ecs.Entity
}

// NewCollect is a helper constructor; you may or may not want to use it.
func NewCollect() *Collect {
    // generate a sine wave for potential item bobbing
    sinSpace := mathutil.SinSpace(0, 2*math.Pi, 3, 60)

    return &Collect{
        itemGravity: 3.0,
        sinSpace:    sinSpace,
        sinSpaceLen: len(sinSpace) - 1,
        itemsToRemove: make([]ecs.Entity, 0),
    }
}

// Init sets up any resources needed by the Collect system.
// If there's no special initialization needed, we can keep it empty.
func (c *Collect) Init() {
    // Potential place for logging or additional data loading
}

// Update is called once per frame to handle item collision, picking them up, and applying gravity.
func (c *Collect) Update() {
    // Query items from ECS that match the filter
    collisionQuery := arc.FilterItem.Query(&kar.WorldECS)

    for collisionQuery.Next() {
        // itemID = item identity?
        // rect = bounding box for collision
        // timers = some struct that includes CollisionCountdown & AnimationIndex
        // durability = how many uses or HP left in the item
        itemID, rect, timers, durability := collisionQuery.Get()

        // Decrement the collision countdown, if any
        if timers.CollisionCountdown > 0 {
            timers.CollisionCountdown--
        } else {
            // Suppose CTRL.Rect is the player bounding box
            if CTRL.Rect.OverlapsRect(rect) {
                // Attempt to add item to player's inventory
                // If success, schedule removing the item from the world
                pickedUp := CTRL.Inventory.AddItemIfEmpty(itemID.ID, durability.Durability)
                if pickedUp {
                    c.itemsToRemove = append(c.itemsToRemove, collisionQuery.Entity())
                }
            }
        }

        // Apply gravity or some vertical movement
        c.applyGravityOrAnimation(rect, timers)
    }

    // Remove any items that were picked up or otherwise flagged for removal
    c.performRemovals()
}

// Draw is a no-op here, but might be used for debug overlays or other visuals.
func (c *Collect) Draw() {
    // e.g., debug draw bounding boxes or item arcs if needed
}

// applyGravityOrAnimation adjusts rect.Y based on collision logic and sinusoidal animation
func (c *Collect) applyGravityOrAnimation(rect *kar.Rect, timers *kar.Timers) {
    // Example: CollideY returns how much it can move down without passing through terrain
    dy := Collider.CollideY(rect.X, rect.Y+8, rect.W, rect.H, c.itemGravity)
    rect.Y += dy

    // Or you could do bobbing instead by using the sinSpace if you prefer:
    // rect.Y += c.sinSpace[timers.AnimationIndex]

    // Increment animation index safely
    timers.AnimationIndex = (timers.AnimationIndex + 1) % c.sinSpaceLen
}

// performRemovals removes scheduled entities from the ECS, then resets the slice.
func (c *Collect) performRemovals() {
    for _, e := range c.itemsToRemove {
        kar.WorldECS.RemoveEntity(e)
    }
    // Clear slice without reallocating
    c.itemsToRemove = c.itemsToRemove[:0]
}
