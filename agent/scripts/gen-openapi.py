import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).parent.parent / "src"))

import yaml  # type: ignore[import-untyped]

from app.app import app


def main() -> None:
    spec = app.openapi()
    print(yaml.dump(spec, sort_keys=False, default_flow_style=False))


if __name__ == "__main__":
    main()
