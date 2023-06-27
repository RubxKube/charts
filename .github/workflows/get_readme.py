#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys, logging, os, yaml
from pathlib import Path
from glob import glob
from yaml.loader import SafeLoader
from jinja2 import Template

def main():
    files = glob('../../**/Chart.yaml', recursive=True)
    charts = []
    for chart in files:
        with open(chart) as f:
            data = yaml.load(f, Loader=SafeLoader)
            print(f"nom : {data['name']} \ndescription: {data['description']}\nversion chart: {data['version']}\nversion app: {data['appVersion']}")
            charts.append([data['name'],data['description'],data['version'], data['appVersion']])
    print(f"Nombre de charts: {len(charts)}")
    table_template=Path('table.j2').read_text()
    tm = Template(table_template)
    tableValue = tm.render({'charts':charts})
    print("----")
    readme_template=Path('./README.md.tmpl').read_text().replace("CHARTS_TABLE",tableValue)
    print(readme_template)
    Path("../../README.md").write_text(readme_template)

if __name__ == "__main__":
    main()
