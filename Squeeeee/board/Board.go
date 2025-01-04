components {
  id: "Board"
  component: "/board/Board.tilemap"
  position {
    x: 160.0
    z: 0.1
  }
}
components {
  id: "Board1"
  component: "/board/Board.script"
}
components {
  id: "score_label"
  component: "/board/score_label.label"
  position {
    x: 80.0
    y: 16.0
    z: 0.1
  }
}
components {
  id: "group_label"
  component: "/board/group_label.label"
  position {
    x: 80.0
    y: 48.0
    z: 0.1
  }
}
components {
  id: "level_label"
  component: "/board/level_label.label"
  position {
    x: 80.0
    y: 80.0
    z: 0.1
  }
}
components {
  id: "restart_button"
  component: "/board/restart_button.gui"
}
embedded_components {
  id: "percentage_label"
  type: "label"
  data: "size {\n"
  "  x: 160.0\n"
  "  y: 32.0\n"
  "}\n"
  "color {\n"
  "  x: 0.502\n"
  "  y: 0.0\n"
  "  z: 0.502\n"
  "}\n"
  "text: \"%\"\n"
  "font: \"/builtins/fonts/default.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  position {
    x: 80.0
    y: 112.0
    z: 0.1
  }
}
embedded_components {
  id: "sfx1"
  type: "sound"
  data: "sound: \"/board/sfx1.wav\"\n"
  ""
}
embedded_components {
  id: "sfx2"
  type: "sound"
  data: "sound: \"/board/sfx2.wav\"\n"
  ""
}
embedded_components {
  id: "sfx3"
  type: "sound"
  data: "sound: \"/board/sfx3.wav\"\n"
  ""
}
embedded_components {
  id: "sfx4"
  type: "sound"
  data: "sound: \"/board/sfx4.wav\"\n"
  ""
}
embedded_components {
  id: "sfx5"
  type: "sound"
  data: "sound: \"/board/sfx5.wav\"\n"
  ""
}
embedded_components {
  id: "sfx6"
  type: "sound"
  data: "sound: \"/board/sfx6.wav\"\n"
  ""
}
embedded_components {
  id: "sfx7"
  type: "sound"
  data: "sound: \"/board/sfx7.wav\"\n"
  ""
}
embedded_components {
  id: "sfx8"
  type: "sound"
  data: "sound: \"/board/sfx8.wav\"\n"
  ""
}
