use super::init::initialize;

fn main() {
    let dialog = initialize();
    dialog.render();
    dialog.refresh();
}
