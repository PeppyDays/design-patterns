use super::{game, magic_maze::MagicMaze, ordinary_maze::OrdinaryMaze};

fn main() {
    let ordinary_maze = OrdinaryMaze::new();
    game::run(ordinary_maze);

    let magic_maze = MagicMaze::new();
    game::run(magic_maze);
}
