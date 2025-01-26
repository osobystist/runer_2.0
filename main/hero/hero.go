components {
  id: "hero"
  component: "/main/hero/hero.script"
}
components {
  id: "hero_moove"
  component: "/main/hero/hero_move.script"
}
embedded_components {
  id: "spinemodel"
  type: "spinemodel"
  data: "spine_scene: \"/main/hero/hero.spinescene\"\n"
  "default_animation: \"run\"\n"
  "skin: \"\"\n"
  "material: \"/defold-spine/assets/spine.material\"\n"
  ""
  scale {
    x: 0.45
    y: 0.45
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"hero\"\n"
  "mask: \"geometry\"\n"
  "mask: \"danger\"\n"
  "mask: \"pickup\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 1.045353\n"
  "      y: 90.93321\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      x: -15.835852\n"
  "      y: 237.93234\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 3\n"
  "    count: 1\n"
  "  }\n"
  "  data: 41.350445\n"
  "  data: 82.146416\n"
  "  data: 10.0\n"
  "  data: 69.1208\n"
  "}\n"
  ""
}
