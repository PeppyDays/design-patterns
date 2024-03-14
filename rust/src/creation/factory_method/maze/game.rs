pub trait Room {
    fn render(&self);
}

pub trait MazeGame {
    type Room: Room;

    fn rooms(&self) -> Vec<Self::Room>;
    fn play(&self) {
        for room in self.rooms() {
            room.render()
        }
    }
}

pub fn run(maze_game: impl MazeGame) {
    println!("Loading resources ..");
    println!("Starting the game ..");

    maze_game.play();
}
