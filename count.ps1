# Initialiser une variable pour le total des lignes
$totalLines = 0

# Parcourir tous les fichiers .go et compter les lignes
Get-ChildItem -Recurse -Filter *.go | ForEach-Object {
    # Compter les lignes (exclure lignes vides et commentaires)
    $lines = (Get-Content $_.FullName | Where-Object { $_ -notmatch "^\s*$" -and $_ -notmatch "^\s*\/\/" }).Count
    # Afficher le fichier et le nombre de lignes
    Write-Output "$($_.FullName): $lines lignes"
    # Ajouter au total
    $totalLines += $lines
}

# Afficher le total
Write-Output "`nTotal des lignes de code Go : $totalLines lignes"
